/**
 * @param {number[]} arr
 * @return {number[]}
 */
var getDistances = function(arr) {
    const intervals = [];

    for (const i in arr) {
        let interval = 0;

        for (const j in arr) {
            interval += arr[i] === arr[j] ? Math.abs(i - j) : 0;
        }

        intervals.push(interval);
    }

    return intervals;
}
