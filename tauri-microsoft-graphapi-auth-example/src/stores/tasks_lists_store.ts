import {derived, writable} from "svelte/store";
import {WellknownListName, type CommandResult, type TaskList} from "../utils/models";
import {invoke} from "@tauri-apps/api/tauri";
import {MessageType, push_new_message} from "./toast_store";

export const task_lists = writable<TaskList[]>([]);
export async function fetch_tasks_lists(refresh_cache: boolean) {
  const function_to_call = refresh_cache ? "get_task_lists" : "get_task_lists_cached";
  const new_task_lists = (await invoke(function_to_call)) as CommandResult<TaskList[]>;
  if (!new_task_lists.success) {
    push_new_message(MessageType.error, new_task_lists.err_message!);
  }
  task_lists.set(new_task_lists.result!);
}

export const sidebar_task_lists = derived(task_lists, (v) =>
  v.filter((s) => s.wellknownListName == WellknownListName.None)
);
