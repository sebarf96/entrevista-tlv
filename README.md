# Entrevista TLV


## Installation

Clonar repo

```bash
git clone https://github.com/sebarf96/entrevista-tlv.git

```

Compilar

```bash
make build
```

Ejecutar

```bash
build/bin/tlv-to-map
```

## Formato de los campos TLV
Cada campo TLV esta compuesto por 3 partes:


 - **Largo**: 2 caracteres que indican el largo del valor, este campo es importante puesto que indica cuantos caracteres leer a continuación.
  - **Tipo**: El tipo tiene un largo de 3 caracteres donde el primer caracter indica el tipo de dato  (A: Alfanumérico y N: Numérico) y dos caracteres para indicar el numero de campo Ejemplo: "01" o "15".
 - **Valor**: Este es el valor del campo, el cual corresponde al valor del campo, su largo esta determinado por la porción **Largo**.

Ejemplo:

> Para "11A05AB398765UJ102N2300" Los campos son:
> - 05 de tipo Alfanumérico de largo 11 y valor AB398765UJ1
> - 23 de tipo Numérico de largo 2 y valor 00
