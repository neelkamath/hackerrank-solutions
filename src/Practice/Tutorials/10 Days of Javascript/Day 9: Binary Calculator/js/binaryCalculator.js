function appendRes(val) {
    document.getElementById('res').innerHTML += val;
}

document.getElementById('btn0').addEventListener('click', () => appendRes('0'));
document.getElementById('btn1').addEventListener('click', () => appendRes('1'));
document.getElementById('btnClr').addEventListener('click', () => document.getElementById('res').innerHTML = '');
document.getElementById('btnSum').addEventListener('click', () => appendRes('+'));
document.getElementById('btnSub').addEventListener('click', () => appendRes('-'));
document.getElementById('btnMul').addEventListener('click', () => appendRes('*'));
document.getElementById('btnDiv').addEventListener('click', () => appendRes('/'));

document.getElementById('btnEql').addEventListener('click', () => {
    let result = document.getElementById('res'),
        nums = result.innerHTML.split(/\D/),
        num1 = parseInt(nums[0], 2),
        num2 = parseInt(nums[1], 2),
        answer;
    switch (/\D/.exec(result.innerHTML)[0]) {
        case '+':
            answer = num1 + num2;
            break;
        case '-':
            answer = num1 - num2;
            break;
        case '*':
            answer = num1 * num2;
            break;
        case '/':
            answer = num1 / num2;
    }
    result.innerHTML = answer.toString(2);
});