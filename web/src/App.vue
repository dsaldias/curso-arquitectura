<template>

  <div> 
    <form @submit="validar_datos">
      <div>
        <label>Nombre:</label>
        <input type="text" v-model="nombre">
      </div>
      
      <div>
        <label>Password:</label>
        <input type="password" v-model="password">
      </div>
      
      <div>
        <button type="submit"> Ingresar </button>
      </div>
      
    </form>
  </div>

</template>

<script setup lang="ts">
import { ref } from 'vue';


const nombre = ref("")
const password = ref("")

const validar_datos = (e:any) => {
  e.preventDefault()

  console.log("usuario:", nombre.value)
  console.log("password:", password.value)
}

</script>
