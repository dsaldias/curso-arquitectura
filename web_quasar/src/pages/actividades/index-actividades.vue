<template>
  <q-page class="q-pa-md">
    <!-- ENCABEZADO -->
    <div class="row items-center justify-between q-mb-md">
      <div>
        <div class="text-h5">Actividades</div>
        <div class="text-grey-7">Gestión de actividades del sistema</div>
      </div>

      <q-btn
        color="primary"
        icon="add"
        label="Nueva actividad"
        @click="nuevaActividad"
      />
    </div>

    <!-- TABLA -->
    <q-table
      title="Lista de actividades"
      :rows="actividades"
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
            @click="editarActividad(props.row)"
          >
            <q-tooltip> Editar actividad </q-tooltip>
          </q-btn>

          <q-btn
            flat
            round
            color="negative"
            icon="delete"
            @click="eliminarActividad(props.row)"
          >
            <q-tooltip> Eliminar actividad </q-tooltip>
          </q-btn>
        </q-td>
      </template>
    </q-table>

    <!-- MODAL -->
    <q-dialog v-model="dialog">
      <q-card style="width: 500px; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">
            {{ editando ? "Editar actividad" : "Nueva actividad" }}
          </div>
        </q-card-section>

        <q-separator />

        <!-- FORMULARIO -->
        <q-form @submit="guardarActividad">
          <q-card-section>
            <q-input
              v-model="form.Nombre"
              label="Nombre"
              outlined
              :rules="[(val) => !!val || 'El nombre es obligatorio']"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Descripcion"
              label="Descripción"
              type="textarea"
              outlined
              :rules="[(val) => !!val || 'La descripción es obligatoria']"
              class="q-mb-md"
            />

            <q-input
              v-model="form.Fecha"
              label="Fecha"
              type="date"
              outlined
              :rules="[(val) => !!val || 'La fecha es obligatoria']"
              class="q-mb-md"
            />
          </q-card-section>

          <q-separator />

          <q-card-actions align="right">
            <q-btn flat label="Cancelar" color="grey" v-close-popup />

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
import { ref, onMounted } from "vue";
import { useQuasar } from "quasar";

const $q = useQuasar();

const actividades = ref([]);
const loading = ref(false);
const guardando = ref(false);

const dialog = ref(false);
const editando = ref(false);

const form = ref({
  Id: "",
  Nombre: "",
  Descripcion: "",
  Fecha: "",
});

const columns = [
  {
    name: "Nombre",
    label: "Nombre",
    field: "Nombre",
    align: "left",
    sortable: true,
  },
  {
    name: "Descripcion",
    label: "Descripción",
    field: "Descripcion",
    align: "left",
    sortable: true,
  },
  {
    name: "Fecha",
    label: "Fecha",
    field: "Fecha",
    align: "left",
    sortable: true,
  },
  {
    name: "acciones",
    label: "Acciones",
    field: "acciones",
    align: "center",
  },
];

// ===============================
// LISTAR ACTIVIDADES
// ===============================

async function cargarActividades() {
  loading.value = true;

  try {
    const response = await fetch("http://localhost:8080/actividades");

    if (!response.ok) {
      throw new Error("Error al obtener las actividades");
    }

    actividades.value = await response.json();
  } catch (error) {
    console.error(error);

    $q.notify({
      type: "negative",
      message: "No se pudieron cargar las actividades",
    });
  } finally {
    loading.value = false;
  }
}

// ===============================
// NUEVA ACTIVIDAD
// ===============================

function nuevaActividad() {
  editando.value = false;

  form.value = {
    Id: "",
    Nombre: "",
    Descripcion: "",
    Fecha: "",
  };

  dialog.value = true;
}

// ===============================
// EDITAR ACTIVIDAD
// ===============================

function editarActividad(actividad) {
  editando.value = true;

  form.value = {
    Id: actividad.Id,
    Nombre: actividad.Nombre,
    Descripcion: actividad.Descripcion,
    Fecha: actividad.Fecha,
  };

  dialog.value = true;
}

// ===============================
// GUARDAR
// ===============================

async function guardarActividad() {
  guardando.value = true;

  try {
    let url = "http://localhost:8080/new-actividad";
    let method = "POST";

    if (editando.value) {
      url = "http://localhost:8080/actualizar-actividad";
      method = "PUT";
    }

    let datitos = JSON.stringify({
      Id: form.value.Id,
      Nombre: form.value.Nombre,
      Descripcion: form.value.Descripcion,
      Fecha: form.value.Fecha,
    });

    console.log(datitos);

    const response = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
      },
      body: datitos,
    });

    if (!response.ok) {
      throw new Error("Error al guardar actividad");
    }

    $q.notify({
      type: "positive",
      message: editando.value
        ? "Actividad actualizada correctamente"
        : "Actividad creada correctamente",
    });

    dialog.value = false;

    await cargarActividades();
  } catch (error) {
    console.error(error);

    $q.notify({
      type: "negative",
      message: "No se pudo guardar la actividad",
    });
  } finally {
    guardando.value = false;
  }
}

// ===============================
// ELIMINAR
// ===============================

function eliminarActividad(actividad) {
  $q.dialog({
    title: "Confirmar",
    message: `¿Deseas eliminar la actividad ${actividad.Nombre}?`,
    cancel: true,
    persistent: true,
  }).onOk(async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/eliminar-actividad?idactividad=` + actividad.Id,
        {
          method: "DELETE",
        },
      );

      if (!response.ok) {
        throw new Error("Error al eliminar actividad");
      }

      $q.notify({
        type: "positive",
        message: "Actividad eliminada correctamente",
      });

      await cargarActividades();
    } catch (error) {
      console.error(error);

      $q.notify({
        type: "negative",
        message: "No se pudo eliminar la actividad",
      });
    }
  });
}

// ===============================
// INICIO
// ===============================

onMounted(() => {
  cargarActividades();
});
</script>
