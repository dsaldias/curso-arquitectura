<template>

  <div> 

    <p>
      Nombre = {{ nombre }}
    </p>

    <p>
      Contrasena = {{ password }}
    </p>

    <p>
      Numero = {{ numero }}
    </p>

    <button @click="cambiar_nombre()"> Cambiar nombre </button>
    <button @click="mostrar_la_clave()"> ver clave </button>
    <button @click="incrementar()"> Incrementar </button>

    <input type="text" v-model="nombre2" placeholder="Nombre...">
    <input type="password" v-model="password2" placeholder="Clave...">
    <input type="number" v-model="numero2">

  </div>

</template>

<script setup lang="ts">
import { ref } from 'vue';

const numero = ref(20)
const nombre = ref("Diego Saldias")
const password = ref("")

const nombre2 = ref("")
const password2 = ref("")
const numero2 = ref(0)

const cambiar_nombre = () => {
  nombre.value = nombre2.value
}

const mostrar_la_clave = () => {
  password.value = password2.value
}

const incrementar = () => {
  numero.value += numero2.value
  
}

</script>
