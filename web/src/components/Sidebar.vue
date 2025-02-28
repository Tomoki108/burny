<template>
    <div class="sidebar">
        <nav>
            <ul>
                <router-link to="/projects" custom v-slot="{ navigate, isActive }">
                    <li :class="{ 'active-li': isActive }" @click="navigate">
                        <font-awesome-icon icon="project-diagram" /> Projects
                    </li>
                </router-link>
                <router-link to="/settings" custom v-slot="{ navigate, isActive }">
                    <li :class="{ 'active-li': isActive }" @click="navigate">
                        <font-awesome-icon icon="cog" /> Settings
                    </li>
                </router-link>
            </ul>
        </nav>
        <div class="signout">
            <router-link to="/sign_in"><font-awesome-icon icon="sign-out-alt" /> Sign out</router-link>
        </div>
    </div>
</template>

<style scoped>
.sidebar {
    width: 200px;
    background-color: var(--color-secondary);
    height: 100vh;
    position: flex;
    top: 0;
    left: 0;
}

li {
    cursor: pointer;
    font-size: var(--font-size-large);
}

li:hover,
div.signout:hover {
    color: var(--color-text-light);
    font-weight: bold;
    background-color: var(--color-tertiary);
}


.active-li {
    color: var(--color-text-light);
    font-weight: bold;
    background-color: var(--color-tertiary);
    border-top: 1px solid var(--color-text-light);
    border-bottom: 1px solid var(--color-text-light);
}

nav {
    margin-top: 20px;
}

nav ul {
    list-style: none;
}

nav ul li {
    padding-top: 20px;
    padding-right: 40px;
    padding-left: 40px;
    padding-bottom: 20px;
}

nav ul li a {
    text-decoration: none;
    color: var(--color-text);
    font-size: var(--font-size-large);
    display: flex;
    align-items: center;
}

nav ul li a svg {
    margin-right: 20px;
}

.signout {
    position: absolute;
    bottom: 20px;
    padding-top: 20px;
    padding-right: 40px;
    padding-left: 40px;
    padding-bottom: 20px;
    width: inherit;
}

.signout a {
    font-weight: 1rem;
    display: flex;
    align-items: center;
    color: inherit;
    font-size: var(--font-size-large);
}

.signout a:hover {
    color: var(--color-text-light);
    font-weight: bold;
}

.signout a svg,
li svg {
    margin-right: 8px;
}
</style>