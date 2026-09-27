
<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <q-page class="row justify-center items-center">

        <q-card style="width: 400px; max-width: 90vw;">
          <q-card-section>
            <div class="text-h5 text-center">
              Iniciar sesión
            </div>
          </q-card-section>

          <q-card-section>
            <q-form @submit.prevent="login">

              <q-input
                v-model="usuario"
                label="Usuario"
                outlined
                class="q-mb-md"
              />

              <q-input
                v-model="password"
                label="Contraseña"
                type="password"
                outlined
                class="q-mb-md"
              />

              <q-btn
                type="submit"
                label="Ingresar"
                color="primary"
                class="full-width"
              />

            </q-form>

            <p>
              {{ respuesta }}
            </p>

          </q-card-section>
        </q-card>

      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter()

const usuario = ref("");
const password = ref("");
const respuesta = ref("");

async function login() {
  respuesta.value = ""

  const response = await fetch("http://localhost:8080/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      usuario1: usuario.value,
      password1: password.value
    })
  });

  const resultado = await response.text();

  console.log("tipo: ",typeof(resultado));


  respuesta.value = resultado

  if (respuesta.value == "Acceso concedido"){
    router.push('/main');
    return;
  }

  console.log("Respuesta del servidor:", resultado);
  console.log(resultado);

  router.push('/main');
}



</script>
```
