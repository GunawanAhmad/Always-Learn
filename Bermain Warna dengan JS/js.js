 const merah = document.querySelector("input[name='merah']");
 const biru = document.querySelector("input[name='biru']");
 const hijau = document.querySelector("input[name='hijau']");

 merah.addEventListener('input', function () {
     const r = merah.value;
     const g = hijau.value;
     const b = biru.value;
     console.log(r);
     document.body.style.backgroundColor = 'rgb(' + r + ',' + g + ',' + b + ')'
 });


 hijau.addEventListener('input', function () {
     const g = hijau.value;
     const r = merah.value;
     const b = biru.value;
     console.log(g);
     document.body.style.backgroundColor = 'rgb(' + r + ',' + g + ',' + b + ')'
 });


 biru.addEventListener('input', function () {
     const b = biru.value;
     const g = hijau.value;
     const r = merah.value;
     console.log(b);
     document.body.style.backgroundColor = 'rgb(' + r + ',' + g + ',' + b + ')'

 });


 // Dibahawah adalh versi yang lebih efektif dan simple (nyontek dari Comment Web Programming Unpas Tentang DOM Video ke 8)

 //  warna.forEach(function (slider) {
 //      slider.addEventListener("input", function () {
 //          const r = document.querySelector("input[name='merah']").value;
 //          const g = document.querySelector("input[name='hijau']").value;
 //          const b = document.querySelector("input[name='biru']").value;
 //          console.log(r, g, b);
 //          document.body.style.backgroundColor = `rgb(${r},${g},${b})`;
 //      })
 //  })


 //Membuat tombol yang dapat mengacak warna
 const acak = document.getElementById('tom')
 acak.addEventListener('click', function () {
     const red = Math.round(Math.random() * 255 + 1);
     const green = Math.round(Math.random() * 255 + 1);
     const blue = Math.round(Math.random() * 255 + 1);
     document.body.style.backgroundColor = 'rgb(' + red + ', ' + green + ', ' + blue + ')';
 });