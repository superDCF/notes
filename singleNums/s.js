function singleNumbers(nums) {
    let k = 0
    nums.forEach(e => {
        k ^= e
    })
    console.log("k", k) // 公异数
    mask = 1

    while ((k & mask) == 0) {
        mask <<= 1
    }
    console.log("mask", mask)
    let a = 0
    let b = 0
    // 这里求公异或数，分为两组，组内自异或运算
    // 两个数异或运算，在分别对这个数作与运算，一定有一个数的结果为0，另一个不为0
    nums.forEach(e => {
        if ((e & mask) == 0) {
            a ^= e
        } else {
            b ^= e
        }
    })
    console.log("ab", [a, b])
    return [a, b]
}
console.log(singleNumbers([4, 2, 4, 6]))