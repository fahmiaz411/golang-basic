class Str {
  constructor(val) {
    this.val = val;
    this.itung = (() => {
      let count = 0;
      for (const i in val) {
        count++;
      }
      return count;
    })();
  }
}

const myName = new Str("hel");
console.log(myName.length);
