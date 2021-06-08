"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const Fila_1 = require("./Fila");
const Tipos_1 = require("./Tipos");
const Tabla_1 = require("./Tabla");
class EtiquetaSimple {
    constructor(nombreTag, listaAtributos, linea, columna) {
        this.nombreTag = nombreTag;
        this.padre = null;
        this.listaAtributos = listaAtributos;
        this.linea = linea;
        this.columna = columna;
        this.listaAtributos.forEach(atributo => {
            atributo.etiquetaContendora = this;
        });
    }
    getName() {
        return this.nombreTag;
    }
    getAmbito() {
        let listaAmbito = [];
        for (let etiqueta = this.padre; etiqueta != null; etiqueta = etiqueta.padre) {
            listaAmbito.push(etiqueta.getName());
        }
        listaAmbito.push("GLOBAL");
        return listaAmbito;
    }
    imprimir() {
        let texto = "";
        texto += "<" + this.nombreTag;
        this.listaAtributos.forEach(atributo => {
            texto += " " + atributo.imprimir();
        });
        texto += "/> \n";
        return texto;
    }
    getAsTable() {
        let tabla = new Tabla_1.Tabla();
        tabla.addFila(new Fila_1.Fila(this.getName(), Tipos_1.Tipos.ETIQUETA_SIMPLE, this.getAmbito(), this.linea, this.columna, this.imprimir()));
        this.listaAtributos.forEach(atributo => {
            tabla.addFila(atributo.getAsRowTable());
        });
        return tabla;
    }
    getErroresSemanticos() {
        let texto = "";
        this.listaAtributos.forEach(atributo => {
            let apariciones = 0;
            for (let atr2 of this.listaAtributos) {
                if (atributo.nombre == atr2.nombre) {
                    apariciones += 1;
                }
                if (apariciones > 1) {
                    texto += `Error(Linea: ${atributo.linea}, Columna: ${atributo.columna}): El atributo '${atributo.nombre}' se encuentra repetido.\n`;
                    break;
                }
            }
        });
        return texto;
    }
}
exports.EtiquetaSimple = EtiquetaSimple;
