<template>
  <q-page class="q-pa-md">

    <!-- ENCABEZADO -->
    <div class="row items-center justify-between q-mb-md">
      <div>
        <div class="text-h5">Usuarios</div>
        <div class="text-grey-7">
          Gestión de usuarios del sistema
        </div>
      </div>

      <q-btn
        color="primary"
        icon="add"
        label="Nuevo usuario"
        @click="nuevoUsuario"
      />
    </div>

    <!-- TABLA -->
    <q-table
      title="Lista de usuarios"
      :rows="usuarios"
      :columns="columns"
      row-key="Id"
      :loading="loading"
      flat
      bordered
    >

      <!-- ACCIONES -->
      <template v-slot:body-cell-acciones="props">
        <q-td :props="props">

          <q-btn
            flat
            round
            color="primary"
            icon="edit"
            @click="editarUsuario(props.row)"
          >
            <q-tooltip>
              Editar usuario
            </q-tooltip>
          </q-btn>

          <q-btn
            flat
            round
            color="negative"
            icon="delete"
            @click="eliminarUsuario(props.row)"
          >
            <q-tooltip>
              Eliminar usuario
            </q-tooltip>
          </q-btn>

        </q-td>
      </template>

    </q-table>

    <!-- MODAL -->
    <q-dialog v-model="dialog">

      <q-card style="width: 500px; max-width: 90vw;">

        <q-card-section>
          <div class="text-h6">
            {{ editando ? 'Editar usuario' : 'Nuevo usuario' }}
          </div>
        </q-card-section>

        <q-separator />

        <!-- FORMULARIO -->
        <q-form @submit="guardarUsuario">

          <q-card-section>

            <q-input
              v-model="form.Nombre"
              label="Nombre"
              outlined
              :rules="[
                val => !!val || 'El nombre es obligatorio'
              ]"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Apellidos"
              label="Apellidos"
              outlined
              :rules="[
                val => !!val || 'Los apellidos son obligatorios'
              ]"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Correo"
              label="Correo electrónico"
              type="email"
              outlined
              :rules="[
                val => !!val || 'El correo es obligatorio',
                val => /.+@.+\..+/.test(val) || 'Correo inválido'
              ]"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Username"
              label="Username"
              outlined
              :rules="[
                val => !!val || 'El username es obligatorio'
              ]"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Password"
              label="Password"
              type="password"
              outlined
              :rules="[
                val => !!val || 'La contraseña es obligatoria'
              ]"
              class="q-mb-md"
            />

          </q-card-section>

          <q-separator />

          <q-card-actions align="right">

            <q-btn
              flat
              label="Cancelar"
              color="grey"
              v-close-popup
            />

            <q-btn
              type="submit"
              label="Guardar"
              color="primary"
              :loading="guardando"
            />

          </q-card-actions>

        </q-form>

      </q-card>

    </q-dialog>

  </q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useQuasar } from 'quasar'

const $q = useQuasar()

const usuarios = ref([])
const loading = ref(false)
const guardando = ref(false)

const dialog = ref(false)
const editando = ref(false)

const form = ref({
  Id: '',
  Nombre: '',
  Apellidos: '',
  Correo: '',
  Username: '',
  Password: ''
})

const columns = [
  {
    name: 'Nombre',
    label: 'Nombre',
    field: 'Nombre',
    align: 'left',
    sortable: true
  },
  {
    name: 'Apellidos',
    label: 'Apellidos',
    field: 'Apellidos',
    align: 'left',
    sortable: true
  },
  {
    name: 'Correo',
    label: 'Correo',
    field: 'Correo',
    align: 'left',
    sortable: true
  },
  {
    name: 'Username',
    label: 'Username',
    field: 'Username',
    align: 'left',
    sortable: true
  },
  {
    name: 'acciones',
    label: 'Acciones',
    field: 'acciones',
    align: 'center'
  }
]


// ===============================
// LISTAR USUARIOS
// ===============================

async function cargarUsuarios () {

  loading.value = true

  try {

    const response = await fetch('http://localhost:8080/usuarios')

    if (!response.ok) {
      throw new Error('Error al obtener los usuarios')
    }

    usuarios.value = await response.json()

  } catch (error) {

    console.error(error)

    $q.notify({
      type: 'negative',
      message: 'No se pudieron cargar los usuarios'
    })

  } finally {

    loading.value = false

  }
}


// ===============================
// NUEVO USUARIO
// ===============================

function nuevoUsuario () {

  editando.value = false

  form.value = {
    Id: '',
    Nombre: '',
    Apellidos: '',
    Correo: '',
    Username: '',
    Password: ''
  }

  dialog.value = true
}


// ===============================
// EDITAR USUARIO
// ===============================

function editarUsuario (usuario) {

  editando.value = true

  form.value = {
    Id: usuario.Id,
    Nombre: usuario.Nombre,
    Apellidos: usuario.Apellidos,
    Correo: usuario.Correo,
    Username: usuario.Username,
    Password: usuario.Password
  }

  dialog.value = true
}


// ===============================
// GUARDAR
// ===============================

async function guardarUsuario () {

  guardando.value = true

  try {

    let url = 'http://localhost:8080/new-usuario'
    let method = 'POST'

    if (editando.value) {
      url = `http://localhost:8080/usuarios/${form.value.Id}`
      method = 'PUT'
    }

    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        Nombre: form.value.Nombre,
        Apellidos: form.value.Apellidos,
        Correo: form.value.Correo,
        Username: form.value.Username,
        Password: form.value.Password
      })
    })

    if (!response.ok) {
      throw new Error('Error al guardar usuario')
    }

    $q.notify({
      type: 'positive',
      message: editando.value
        ? 'Usuario actualizado correctamente'
        : 'Usuario creado correctamente'
    })

    dialog.value = false

    await cargarUsuarios()

  } catch (error) {

    console.error(error)

    $q.notify({
      type: 'negative',
      message: 'No se pudo guardar el usuario'
    })

  } finally {

    guardando.value = false

  }
}


// ===============================
// ELIMINAR
// ===============================

function eliminarUsuario (usuario) {

  $q.dialog({
    title: 'Confirmar',
    message: `¿Deseas eliminar al usuario ${usuario.Username}?`,
    cancel: true,
    persistent: true
  }).onOk(async () => {

    try {

      const response = await fetch(
        `http://localhost:8080/usuarios/${usuario.Id}`,
        {
          method: 'DELETE'
        }
      )

      if (!response.ok) {
        throw new Error('Error al eliminar usuario')
      }

      $q.notify({
        type: 'positive',
        message: 'Usuario eliminado correctamente'
      })

      await cargarUsuarios()

    } catch (error) {

      console.error(error)

      $q.notify({
        type: 'negative',
        message: 'No se pudo eliminar el usuario'
      })

    }

  })
}


// ===============================
// INICIO
// ===============================

onMounted(() => {
  cargarUsuarios()
})
</script>
