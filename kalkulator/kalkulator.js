const angka = document.querySelectorAll('.tombol .angka')
const layar = document.querySelector('#proses')

function cari1() {
    for (nama of angka) {
        nama.addEventListener('click', function (e) {
            var tombolClick = e.target;
            var nilaiTombol = tombolClick.innerText;
            if (nilaiTombol == 'C') {
                layar.innerText = ''
            } else if (nilaiTombol == 'Del') {
                layar.innerText = layar.innerText.slice(0, -1);
            } else if (nilaiTombol == '=') {
                layar.innerText = eval(layar.innerText);
            } else {
                layar.innerText = layar.innerText + nilaiTombol;
            }

        });
    }

}
cari1()