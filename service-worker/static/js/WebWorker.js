console.log(`%c Started Web Worker`, 'font-size: 20px; color: #00ff00; text-align: center; margin-left: calc(20%); padding: 5px 0 15px 0;');
let [num, limit] = [0,0];

self.addEventListener('message', e => {
  limit = e.data || 1000;
  const calArmstrong = number => {
    let originalNumber, remainder
    let [result, n] = [0, 0];
    originalNumber = number;
    while (originalNumber !== 0) {
      originalNumber = parseInt(originalNumber / 10);
      ++n;
    }
    originalNumber = number;
    while (originalNumber !== 0) {
      remainder = originalNumber % 10;
      result += Math.pow(remainder, n);
      originalNumber = parseInt(originalNumber / 10);
    }
    (result === number) && postMessage(number);
  }
  
  while (num < limit) {
    console.log(limit);
    calArmstrong(++num);
  }
  // self.close()
}, false);
