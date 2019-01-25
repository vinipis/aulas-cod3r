//função em JS é First-Class Object (Citizens)
//higher-order function

// criar de forma literal
function fun1() { }

// armazenar em uma variavel
const fun2 = function() { }

//Armazenar em um arry
const arry = [function (a, b) { return a +b} , fun1, fun2]
console.log(arry[0](2, 3))

//Armazenar em um atributo de objetos
const obj = { }
obj.falar = function() { return 'Opa'}
console.log(obj.falar())

//Passar funções como parametro
function run(fun) {
    fun()
}
run(function() { console.log('Executando...') })

//Uma função pode retornar/conter uma função
function soma(a, b) {
    return function (c) {
        console.log(a + b+ c)
    }
}
soma(2, 3)(4)
const cincoMais = soma(2, 3)
cincoMais(4)