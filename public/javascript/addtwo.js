
function add_two() {
    let t = document.getElementById("inputtarget").value;
    let arr = document.getElementById("inputarray").value;
    let parts = arr.split(",");
    let ind = {}
    let res = "";
    for (let i in parts) {
        parts[i] = parts[i].trim();
        let df = t - parts[i];
        if (df in ind) {
            res = ind[df] + ", " + i;
            break
        } else {
            ind[parts[i]] = i
        }
    }
    document.getElementById("output").innerText = res;
}