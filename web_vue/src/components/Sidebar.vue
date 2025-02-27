<template>
    <div class="sidebar">
        <nav>
            <ul>
                <li><router-link to="/home"><i class="icon-home"></i> Home</router-link></li>
                <li><router-link to="/projects"><i class="icon-projects"></i> Projects</router-link></li>
                <li><router-link to="/settings"><i class="icon-settings"></i> Settings</router-link></li>
            </ul>
        </nav>
        <div class="logout">
            <router-link to="/login"><i class="icon-logout"></i> Log out</router-link>
        </div>
    </div>
</template>

<style scoped>
.sidebar {
    width: 200px;
    background-color: #f0f0f0;
    height: 100vh;
    padding: 20px;
}

nav ul {
    list-style: none;
    padding: 0;
}

nav ul li {
    margin-bottom: 20px;
}

nav ul li a {
    text-decoration: none;
    color: #000;
    display: flex;
    align-items: center;
}

nav ul li a i {
    margin-right: 10px;
}

.logout {
    position: absolute;
    bottom: 20px;
    width: 100%;
}

.logout a {
    text-decoration: none;
    color: #000;
    display: flex;
    align-items: center;
}

.logout a i {
    margin-right: 10px;
}
</style>