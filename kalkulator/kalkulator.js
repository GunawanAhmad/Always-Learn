const angka = document.querySelectorAll('.tombol .angka')
const layar = document.querySelector('#proses')

function cari1() {
    for (nama of angka) {
        nama.addEventListener('click', function (e) {
            var tombolClick = e.target;
            var nilaiTombol = tombolClick.innerText;
            if (nilaiTombol == 'C') {
                layar.value = ''
            } else if (nilaiTombol == 'Del') {
                layar.value = layar.value.slice(0, 1);
            } else if (nilaiTombol == '=') {
                layar.value = eval(layar.value);
            } else {
                layar.value = layar.value + nilaiTombol;
            }

        });
    }

}
cari1()