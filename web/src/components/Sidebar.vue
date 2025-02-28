<template>
    <div class="sidebar">
        <nav>
            <ul>
                <li><router-link to="/projects"><font-awesome-icon icon="project-diagram" /> Projects</router-link></li>
                <li><router-link to="/settings"><font-awesome-icon icon="cog" /> Settings</router-link></li>
            </ul>
        </nav>
        <div class="logout">
            <router-link to="/login"><font-awesome-icon icon="sign-out-alt" /> Log out</router-link>
        </div>
    </div>
</template>

<style scoped>
.sidebar {
    width: 200px;
    background-color: var(--color-secondary);
    height: 100vh;
    /* padding: 40px; */
    position: flex;
    top: 0;
    left: 0;
}

.router-link-active,
.router-link-exact-active {
    color: var(--color-primary);
    font-weight: bold;
    /* Alternatively, add background color or any other styling */
}

nav {
    margin-top: 40px;
}

nav ul {
    list-style: none;
    padding: 0;
}

nav ul li {
    margin-left: 40px;
    margin-right: 40px;
    margin-bottom: 20px;
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

.logout {
    position: absolute;
    margin-left: 40px;
    bottom: 40px;
    width: 100%;
}

.logout a {
    text-decoration: none;
    color: var(--color-text);
    display: flex;
    align-items: center;
}

.logout a svg {
    margin-right: 20px;
}
</style>