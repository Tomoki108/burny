<template>
    <div class="home-container">
        <header class="hero">
            <h1 class="hero-title">Burny</h1>
            <p class="hero-subtitle">Simplify Project Management with Burn-up Charts</p>
            <div class="hero-actions">
                <router-link to="/sign_in" class="button primary">Sign In</router-link>
                <button @click="scrollToAbout" class="button secondary">Learn More</button>
            </div>
        </header>

        <section id="about" class="about-section">
            <div class="content-container">
                <h2>About Burny</h2>
                <div class="about-content">
                    <div class="about-text">
                        <p>
                            Burny is a burn-up chart tool designed to visually track project management.
                            Easily monitor progress and visualize project completion predictions.
                        </p>
                        <p>
                            Developed to improve team performance and enhance project transparency.
                            Complex project management made simple with an intuitive interface.
                        </p>
                    </div>
                    <div class="about-image">
                        <img src="../assets/dog_emoji.png" alt="Burny Logo" />
                    </div>
                </div>
            </div>
        </section>

        <section class="features-section">
            <div class="content-container">
                <h2>Key Features</h2>
                <div class="features-grid">
                    <div class="feature-card">
                        <h3>Burn-up Charts</h3>
                        <p>Visually display project progress and easily verify completion predictions.</p>
                    </div>
                    <div class="feature-card">
                        <h3>Sprint Management</h3>
                        <p>Effectively manage agile development sprints and improve team productivity.</p>
                    </div>
                    <div class="feature-card">
                        <h3>Project Tracking</h3>
                        <p>Centrally manage multiple projects and check progress at a glance.</p>
                    </div>
                </div>
            </div>
        </section>

        <section class="why-section">
            <div class="content-container">
                <h2>Why We Built This</h2>
                <p>
                    This project began with the desire to manage team progress more efficiently.
                    By providing features missing in existing tools and a more intuitive interface,
                    we aim to reduce the burden of project management and allow teams to focus on higher-value work.
                </p>
            </div>
        </section>

        <section class="cta-section">
            <div class="content-container">
                <h2>Get Started Now</h2>
                <div class="cta-buttons">
                    <router-link to="/sign_in" class="button primary">Sign In</router-link>
                    <button @click="goToSignUp" class="button secondary">Sign Up</button>
                </div>
            </div>
        </section>

        <footer class="site-footer">
            <div class="content-container">
                <div class="footer-links">
                    <a href="https://github.com/yourusername/burny" target="_blank" rel="noopener noreferrer"
                        class="github-link">
                        <span>View Source Code on GitHub</span>
                    </a>
                </div>
                <p class="copyright">&copy; {{ new Date().getFullYear() }} Burny. All rights reserved.</p>
            </div>
        </footer>
    </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { PATH_SIGN_IN } from '../router';

const router = useRouter();

const scrollToAbout = () => {
    document.getElementById('about')?.scrollIntoView({ behavior: 'smooth' });
};

const goToSignUp = () => {
    router.push({ path: PATH_SIGN_IN, query: { signup: 'true' } });
};
</script>

<style scoped>
.home-container {
    width: 100%;
    min-height: 100vh;
    background-color: var(--color-background);
    color: var(--color-text);
}

.hero {
    height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    background: linear-gradient(var(--color-tertiary), var(--color-secondary));
    color: var(--color-text-light);
    padding: 2rem;
}

.hero-title {
    font-size: 4rem;
    margin-bottom: 1rem;
}

.hero-subtitle {
    font-size: 1.5rem;
    margin-bottom: 2rem;
    max-width: 600px;
}

.hero-actions {
    display: flex;
    gap: 1rem;
}

.content-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 2rem;
}

.about-section,
.features-section,
.why-section,
.cta-section {
    padding: 5rem 0;
}

.about-section {
    background-color: var(--color-background);
}

.about-content {
    display: flex;
    align-items: center;
    gap: 2rem;
}

.about-text {
    flex: 1;
}

.about-image {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
}

.about-image img {
    max-width: 300px;
    height: auto;
}

.features-section {
    background-color: var(--color-background-alt);
}

.features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-top: 2rem;
}

.feature-card {
    background-color: var(--color-background);
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.why-section {
    background-color: var(--color-background);
}

.cta-section {
    background-color: var(--color-tertiary);
    color: var(--color-text-light);
    text-align: center;
}

.cta-buttons {
    display: flex;
    justify-content: center;
    gap: 1rem;
    margin-top: 2rem;
}

.site-footer {
    background-color: var(--color-background-dark);
    color: var(--color-text-light);
    padding: 2rem 0;
}

.footer-links {
    display: flex;
    justify-content: center;
    margin-bottom: 1rem;
}

.github-link {
    color: var(--color-text-light);
    text-decoration: none;
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.github-link:hover {
    text-decoration: underline;
}

.copyright {
    text-align: center;
    font-size: 0.875rem;
    margin-top: 1rem;
}

.button {
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 4px;
    font-weight: bold;
    cursor: pointer;
    text-decoration: none;
    transition: background-color 0.3s;
}

.button.primary {
    background-color: var(--color-primary);
    color: var(--color-text-light);
}

.button.secondary {
    background-color: transparent;
    color: var(--color-text-light);
    border: 2px solid var(--color-text-light);
}

.button.primary:hover {
    background-color: var(--color-primary-dark);
}

.button.secondary:hover {
    background-color: rgba(255, 255, 255, 0.1);
}

@media (max-width: 768px) {
    .about-content {
        flex-direction: column;
    }

    .hero-title {
        font-size: 3rem;
    }

    .hero-subtitle {
        font-size: 1.25rem;
    }
}
</style>