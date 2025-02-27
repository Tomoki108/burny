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
    background-color: #f0f0f0;
    height: 100vh;
    padding: 20px;
    position: fixed;
    /* サイドバーを左側に固定 */
    top: 0;
    left: 0;
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

nav ul li a svg {
    margin-right: 20px;
    /* アイコンとテキストの間に隙間を追加 */
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

.logout a svg {
    margin-right: 20px;
    /* アイコンとテキストの間に隙間を追加 */
}
</style>