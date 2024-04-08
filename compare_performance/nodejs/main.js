async function asyncFunction() {
  console.log("Async function started");
  await new Promise((resolve) => setTimeout(resolve, 2000));
  console.log("Async function completed");
}

function syncFunction() {
  console.log("Sync function started");
  console.log("Sync function completed");
}

asyncFunction();
syncFunction();

// nodejs is non blocking and single threaded.
// good for IO bound tasks because of non blocking nature
// not so good for CPU bound tasks because of being single threaded. 