/**
 * @param {number} startValue
 * @param {number} target
 * @return {number}
 */

var brokenCalc = function(startValue, target) {
    let ans;

    for (ans = 0; target > startValue; target /= 2, ans++) {
        if (target % 2 !== 0) {
            target++;
            ans++;
        }
    }

    ans += startValue - target;

    return ans;
}
