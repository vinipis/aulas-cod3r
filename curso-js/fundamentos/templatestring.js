const nome = 'rebeca'
const concatenando = 'olá ' + nome + "!"
const template = `
    Ola
    ${nome}
    !`
console.log(concatenando, template)

//expressoes
console.log(`1 + 1 = ${1 + 1}`)

const up = texto => texto.toUpperCase()
console.log(`ei.. ${up('cuidado')}!`)