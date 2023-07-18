import { createWebHistory, createRouter } from "vue-router";
import User from "@/components/User.vue"
import User2 from "@/components/User2.vue"


const routes = [
    { path : "/user", name : "User", component : User },
    { path : "/user2", name : "User2", component : User2 },
]

const router = createRouter({
    history : createWebHistory(),
    routes : routes
});

export default router;
