<template>
  <q-page class="q-pa-md">
    <!-- ENCABEZADO -->
    <div class="row items-center justify-between q-mb-md">
      <div>
        <div class="text-h5">Tareas</div>
        <div class="text-grey-7">Gestión de tareas del sistema</div>
      </div>

      <q-btn
        color="primary"
        icon="add"
        label="Nueva tarea"
        @click="nuevaTarea"
      />
    </div>

    <!-- TABLA -->
    <q-table
      title="Lista de tareas"
      :rows="tareas"
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
            @click="editarTarea(props.row)"
          >
            <q-tooltip> Editar tarea </q-tooltip>
          </q-btn>

          <q-btn
            flat
            round
            color="negative"
            icon="delete"
            @click="eliminarTarea(props.row)"
          >
            <q-tooltip> Eliminar tarea </q-tooltip>
          </q-btn>
        </q-td>
      </template>
    </q-table>

    <!-- MODAL -->
    <q-dialog v-model="dialog">
      <q-card style="width: 500px; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">
            {{ editando ? "Editar tarea" : "Nueva tarea" }}
          </div>
        </q-card-section>

        <q-separator />

        <!-- FORMULARIO -->
        <q-form @submit="guardarTarea">
          <q-card-section>
            <q-input
              v-model="form.Nombre"
              label="Nombre"
              outlined
              :rules="[(val) => !!val || 'El nombre es obligatorio']"
              class="q-mb-md"
            />

            <q-input
              v-model="form.ActividadId"
              label="ID actividad"
              type="number"
              outlined
              :rules="[(val) => !!val || 'La actividad es obligatoria']"
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

const tareas = ref([]);
const loading = ref(false);
const guardando = ref(false);

const dialog = ref(false);
const editando = ref(false);

const form = ref({
  Id: "",
  Nombre: "",
  ActividadId: "",
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
    name: "ActividadId",
    label: "ID actividad",
    field: "ActividadId",
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
// LISTAR TAREAS
// ===============================

async function cargarTareas() {
  loading.value = true;

  try {
    const response = await fetch("http://localhost:8080/tareas");

    if (!response.ok) {
      throw new Error("Error al obtener las tareas");
    }

    tareas.value = await response.json();
  } catch (error) {
    console.error(error);

    $q.notify({
      type: "negative",
      message: "No se pudieron cargar las tareas",
    });
  } finally {
    loading.value = false;
  }
}

// ===============================
// NUEVA TAREA
// ===============================

function nuevaTarea() {
  editando.value = false;

  form.value = {
    Id: "",
    Nombre: "",
    ActividadId: "",
  };

  dialog.value = true;
}

// ===============================
// EDITAR TAREA
// ===============================

function editarTarea(tarea) {
  editando.value = true;

  form.value = {
    Id: tarea.Id,
    Nombre: tarea.Nombre,
    ActividadId: tarea.ActividadId,
  };

  dialog.value = true;
}

// ===============================
// GUARDAR
// ===============================

async function guardarTarea() {
  guardando.value = true;

  try {
    let url = "http://localhost:8080/new-tarea";
    let method = "POST";

    if (editando.value) {
      url = "http://localhost:8080/actualizar-tarea";
      method = "PUT";
    }

    let datitos = JSON.stringify({
      Id: form.value.Id,
      Nombre: form.value.Nombre,
      ActividadId: Number(form.value.ActividadId),
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
      throw new Error("Error al guardar tarea");
    }

    $q.notify({
      type: "positive",
      message: editando.value
        ? "Tarea actualizada correctamente"
        : "Tarea creada correctamente",
    });

    dialog.value = false;

    await cargarTareas();
  } catch (error) {
    console.error(error);

    $q.notify({
      type: "negative",
      message: "No se pudo guardar la tarea",
    });
  } finally {
    guardando.value = false;
  }
}

// ===============================
// ELIMINAR
// ===============================

function eliminarTarea(tarea) {
  $q.dialog({
    title: "Confirmar",
    message: `¿Deseas eliminar la tarea ${tarea.Nombre}?`,
    cancel: true,
    persistent: true,
  }).onOk(async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/eliminar-tarea?idtarea=` + tarea.Id,
        {
          method: "DELETE",
        },
      );

      if (!response.ok) {
        throw new Error("Error al eliminar tarea");
      }

      $q.notify({
        type: "positive",
        message: "Tarea eliminada correctamente",
      });

      await cargarTareas();
    } catch (error) {
      console.error(error);

      $q.notify({
        type: "negative",
        message: "No se pudo eliminar la tarea",
      });
    }
  });
}

// ===============================
// INICIO
// ===============================

onMounted(() => {
  cargarTareas();
});
</script>
