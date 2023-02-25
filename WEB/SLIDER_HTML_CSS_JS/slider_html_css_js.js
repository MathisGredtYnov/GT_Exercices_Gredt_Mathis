const slide = ["ASSETS/Nepal.jpg", "ASSETS/Numbani.jpg", "ASSETS/Rialto.jpg", "ASSETS/King's_Row.jpg"];
let numero = 0;

function ChangeSlide(sens) {
    numero = numero + sens;
    if (numero < 0)
        numero = slide.length - 1;
    if (numero > slide.length - 1)
        numero = 0;
    document.getElementById("slide").src = slide[numero];
}