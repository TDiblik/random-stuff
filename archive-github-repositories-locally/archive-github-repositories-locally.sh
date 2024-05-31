#!/bin/bash

# Constants
NUMBER_OF_DAYS_TO_KEEP=90 # Feel free to set this to 99999999 if you want to retain your filed forever

# Init archived_repos dir
mkdir archived_repos;
cd archived_repos;

# If the repo gets delete / renamed / whatever, it also vanishes from the list returned by the API, 
# meaning that it will not be affected by future program executions
repo_names_to_clone=($(curl --request GET --url https://api.github.com/user/repos --header "authorization: Bearer $GITHUB_API_TOKEN" --header "content-type: application/json" | jq ".[].full_name"))
for repo_name_raw in "${repo_names_to_clone[@]}"; do 
    repo_name=$(echo $repo_name_raw | cut -c2- | rev | cut -c2- | rev);
    mkdir -p $repo_name;
    cd $repo_name;
    
    # Remove oldest file (if it classifies for deletion based on NUMBER_OF_DAYS_TO_KEEP variable)
    tar_files=$(find . -maxdepth 1 -type f -printf "%f\n" | grep .tar.bz2);
    sorted_dirs=$(printf '%s\n' "${tar_files[@]}" | sort);
    file_to_delete=($sorted_dirs);
    file_to_delete_formatted=$(echo $file_to_delete | rev | cut -c9- | rev);
    let latest_file_date_diff=(`date +%s`-`date +%s -d $file_to_delete_formatted`)/86400;

    if (($latest_file_date_diff >= $NUMBER_OF_DAYS_TO_KEEP)); then
        rm -rf $file_to_delete;
    fi
    
    # Clone repo
    rm -rf temp
    mkdir temp
    cd temp
    repo_url=https://$GITHUB_API_TOKEN@github.com/$repo_name.git
    echo "---- Cloning $repo_url ----";
    git clone $repo_url;

    # Make archive from that repo
    echo "---- Making tar archive $repo_url ----"
    today=$(date +"%Y-%m-%d");
    new_tar_file_name="$today.tar.bz2"
    cloned_dir=$(ls | grep -vw -e .);
    tar -cjf $new_tar_file_name $cloned_dir
    cd ..
    mv temp/$new_tar_file_name .
    rm -rf temp
    cd ../../
done;
