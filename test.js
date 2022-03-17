// const a = [1, 2, 3];
// const b = [4, 5, 6];

// function zip(a, b) {
//   const max = Math.max(a.length, b.length);
//   const result = [];
//   for (let i = 0; i < max; i++) {
//     result.push(a[i], b[i]);
//   }
//   return result;
// }

// console.log(zip(a, b));

// class Form {
//   #nama;
//   #uts;
//   #uas;
//   #tugas;
//   constructor(props) {
//     this.nim = props.nim;
//     this.#nama = props.nama;
//     this.#uts = props.uts;
//     this.#uas = props.uas;
//     this.#tugas = props.tugas;
//   }

//   get graduate() {
//     let statusNilai =
//       (30 / 100) * this.#uts +
//       (40 / 100) * this.#uas +
//       (30 / 100) * this.#tugas;
//     return statusNilai >= 60 ? "LULUS" : "TIDAK LULUS";
//   }
// }

// const form = new Form({
//   nim: "00112233",
//   nama: "fahmi",
//   uts: 0,
//   uas: 0,
//   tugas: 100,
// });

// function range(s, e) {
//   const arr = [];
//   if (!s || !e) {
//     console.log(s, e);
//     return [];
//   } else {
//     if (s < e) {
//       for (let i = s; i <= e; i++) {
//         console.log(i < e);
//         arr.push(i);
//       }
//     } else {
//       for (let i = s; i >= e; i--) {
//         arr.push(i);
//       }
//     }
//   }
//   return arr;
// }

// function run() {
//   const args = process.argv;
//   let hasil = range(+args[3], +args[4]);
//   console.log(hasil);
// }

// run();

function group(array, keys = [Object.keys(array[0])[0]]) {
  return array
    .filter((data, index, arr) => {
      let state = false;
      const array_keys = (key) => arr.map((d) => d[key]);
      for (const i in keys) {
        const key_index = array_keys(keys[i]).indexOf(data[keys[i]]);
        state = key_index == index;
        // if (data.id == 2) console.log(state, key_index, index);
        if (state) return state;

        for (const j in keys) {
          const key_index = array_keys(keys[j]).lastIndexOf(data[keys[j]]);
          // state = key_index == index;

          // if (state) return state;

          const k1 = keys[i],
            k2 = keys[j];
          // if (k1 != k2) {
          state =
            data[k1] != arr[key_index][k1] || data[k2] != arr[key_index][k2];
          if (data.id == 2 || data.id == 3)
            console.log(
              data.id,
              state,
              data[k1],
              arr[key_index][k1],
              key_index,
              i
            );
          if (state) return state;
          // }
        }
        // console.log("break I\n");
      }
      // console.log("break MAP\n");
      return state;
    })
    .map((datam) => {
      const obj = {};
      for (const key of keys) {
        obj["main_" + key] = datam[key];
      }
      obj["lists"] = array.filter((dataf) => {
        let state = true;
        for (const key of keys) {
          state = datam[key] == dataf[key];
          if (!state) return state;
        }
        return state;
      });

      return obj;
    });
}

const arr = {
  data: [
    {
      id: 9,
      no_order: 12121312132,
      user_id: 2,
      symbol_id: 7,
      symbol_name: "BTC",
      asset_id: 1,
      price: 12.16,
      buy: 133.76,
      coin: 11,
      date: "2022-03-17",
    },
    {
      id: 6,
      no_order: 12121356,
      user_id: 2,
      symbol_id: 7,
      symbol_name: "BTC",
      asset_id: 1,
      price: 8,
      buy: 88,
      coin: 11,
      date: "2022-03-17",
    },
    {
      id: 8,
      no_order: 121213788,
      user_id: 2,
      symbol_id: 6,
      symbol_name: "ETH",
      asset_id: 1,
      price: 11.5,
      buy: 132.25,
      coin: 11.5,
      date: "2022-03-17",
    },
    {
      id: 7,
      no_order: 121213777,
      user_id: 2,
      symbol_id: 6,
      symbol_name: "ETH",
      asset_id: 1,
      price: 12,
      buy: 132,
      coin: 11,
      date: "2022-03-17",
    },
    {
      id: 5,
      no_order: 14322432,
      user_id: 2,
      symbol_id: 7,
      symbol_name: "BTC",
      asset_id: 1,
      price: 14,
      buy: 140,
      coin: 10,
      date: "2022-03-16",
    },
    {
      id: 4,
      no_order: 24321424,
      user_id: 2,
      symbol_id: 7,
      symbol_name: "BTC",
      asset_id: 1,
      price: 13,
      buy: 130,
      coin: 10,
      date: "2022-03-16",
    },
    {
      id: 3,
      no_order: 121218,
      user_id: 2,
      symbol_id: 6,
      symbol_name: "ETH",
      asset_id: 1,
      price: 12,
      buy: 131,
      coin: 11,
      date: "2022-03-16",
    },
    {
      id: 2,
      no_order: 121215,
      user_id: 2,
      symbol_id: 6,
      symbol_name: "ETH",
      asset_id: 1,
      price: 12,
      buy: 135,
      coin: 11,
      date: "2022-03-16",
    },
    {
      id: 1,
      no_order: 121213,
      user_id: 2,
      symbol_id: 6,
      symbol_name: "ETH",
      asset_id: 1,
      price: 12,
      buy: 132,
      coin: 11,
      date: "2022-03-15",
    },
  ],
};

console.log(group(arr.data, ["date", "symbol_name"]));
