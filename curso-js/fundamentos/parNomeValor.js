//Par mome/valor
const saudacao = 'Opa' //contexto léxico 1 contexto no qual sua variavel foi definida no codigo

function exec() {
    const saudacao = 'falaaa' // contexto léxico 2
    return saudacao
}

//Objetos são grupos aninhados de pares nome/valor
const cliente = {
    nome: 'Pedro',
    idade: 32,
    peso: 90,
    endereco: {
        logadouro: 'Rua muito legal',
        numero: 123
    }
}

console.log(saudacao)
console.log(exec())
console.log(cliente)