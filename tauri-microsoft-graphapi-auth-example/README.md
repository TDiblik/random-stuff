# Abandoned!

This was originally supposed to be a Microsoft To Do app, however I didn't have enought time to actually implement it beyond Azure AD login and some simple Graph API queries. (also, the Graph API endpoints for To Do are mostly still in beta, for some dumb reason, and I don't want to waste my time just for Microsoft to decide that they want to totally change the API)

I'm putting it here cuz I feel like it could be a really nice 'template' for anybody who wants to use Graph API with Rust/Tauri.

If you ever want to start it up, the "credentials" (which should be public, so no security issue there xd) inside the `./src-tauri/src/config.rs` are already invalidated, so you'll need to create your own (the Azure AD setttings are screenshoted inside the `./readme` folder).
