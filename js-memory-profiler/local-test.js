
const mainModule = require("./index.js")

const snap1 = mainModule.getHeapSnapshot(heavyTask)

function heavyTask(){
  console.log("1e6=", 1e6);

  const bigArr = [];
  for (let i = 0; i < 1e6; i++){
    bigArr.push({i, content:"hi " + i})
  }
}

const loggerMemUsage = mainModule.logMemUsageEach(500)
