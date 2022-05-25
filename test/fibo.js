
function fibo(n) {
    if(n===1 || n===2) {
        return 1;
    }    
    return fibo(n-1) + fibo(n-2);
}

for(let i=1;i<5;i++) {
    console.log(fibo(i));
}