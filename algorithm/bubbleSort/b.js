function BubbleSort(arr) {
    if (!(arr instanceof Array)) return "arr is not array"
    for (let i = 0; i < arr.length - 1; i++) {
        for (let j = 0; j < arr.length - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                const tmp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = tmp
            }
        }
    }
    return arr
}

let arr = [2, 4, 6, 5, 3, 1, 100]

console.log(BubbleSort(arr))
