
/**
 * Heap Snapshot Capture – Take snapshots of memory usage at different points in time.
 */

function logMemUsageEach(timer){
    const getMemUsage = process.memoryUsage;
  
    setInterval(() => {
        console.log("Mem usage at " + new Date().toString() + " is:", process.memoryUsage())
      }, timer
    )
  }
  
  function getHeapSnapshot(task){
    console.log("process.release=",process.release)
  
    // browser engine
    if(typeof document !== "undefined"){
      return ;
    }
  
    // nodejs engine
    if(typeof process !== 'undefined' && process.release.name === 'node'){
      console.log("Mem before: ", process.memoryUsage())
      task();
      console.log("Mem after: ", process.memoryUsage())
  
      console.log("Unit: bytes")
      console.log("rss: explain how much mem the OS assign for This process")
      console.log("heapTotal: used and free spaces, heap auto grow as needed")
      console.log("heapUsed: heap actually in use")
      console.log("external: mem managed outside V8, like using by C++ bindings, native modules")
      console.log("arrayBuffers: mem use for binary data")
  
      return process.memoryUsage;
    }
  
    throw new Error("The JS Engine not found")
  }
  
  module.exports = { getHeapSnapshot, logMemUsageEach }
  