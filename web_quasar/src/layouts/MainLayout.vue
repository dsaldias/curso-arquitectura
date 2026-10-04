<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title> Gestor de tareas </q-toolbar-title>

        <div> version 1 </div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header> Menu lateral </q-item-label>

        <EssentialLink
          v-for="link in linksList"
          :key="link.label"
          v-bind="link"
        />

        <!-- <q-item clickable tag="a" target="_blank" href="menu_2">
          <q-item-section avatar>
            <q-icon name="home" />
          </q-item-section>

          <q-item-section>
            <q-item-label> Inicio  </q-item-label>
            <q-item-label caption> Bienvenidos </q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable tag="a" target="_blank" href="menu_2">
          <q-item-section avatar>
            <q-icon name="person" />
          </q-item-section>

          <q-item-section>
            <q-item-label> Usuarios  </q-item-label>
            <q-item-label caption> gestion de usuarios </q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable tag="a" target="_blank" href="menu_2">
          <q-item-section avatar>
            <q-icon name="person" />
          </q-item-section>

          <q-item-section>
            <q-item-label> Roles  </q-item-label>
            <q-item-label caption> gestion de usuarios </q-item-label>
          </q-item-section>
        </q-item> -->

      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from "vue";
import EssentialLink, {
  type EssentialLinkProps
} from "@/components/EssentialLink.vue";

const linksList: EssentialLinkProps[] = [
  {
    label: "Usuarios",
    caption: "gestion de usuarios",
    icon: "school",
    link: "main/lista-usuarios"
  },
  {
    label: "Actividades",
    caption: "creacion de actividades",
    icon: "code",
    link: "https://github.com/quasarframework"
  },
  {
    label: "Tareas",
    caption: "lista de tareas",
    icon: "chat",
    link: "https://chat.quasar.dev"
  },
];

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
