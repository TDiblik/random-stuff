function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function run(delay_ms = 0) {
  let event = new MouseEvent("mousedown", { bubbles: true, cancelable: true, view: window });
  for (let i = 0; i < 31; i++) {
    document.querySelector('[data-aim-target="true"]').dispatchEvent(event);
    if (delay_ms != 0) {
      await sleep(delay_ms);
    }
  }
}
