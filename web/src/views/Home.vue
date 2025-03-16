<template>
    <div class="home-container">
        <v-card class="px-2 py-2 text-center">⚠️ This application is in beta version and might be discontinued or
            changed without notice. ⚠️</v-card>
        <header class="hero">
            <h1 class="hero-title">Burny 🐶</h1>
            <p class="hero-subtitle">Simplify Project Management with Burn-up Charts</p>
            <div class="hero-actions">
                <button @click="openSignInModal(false)" class="button-large">Sign In</button>
                <button @click="scrollToAbout" class="button-large nav_link">Learn More</button>
            </div>

        </header>

        <section id="about" class="about-section">
            <div class="content-container">
                <h2>About Burny</h2>
                <div class="about-content">
                    <div class="about-text">
                        <p>
                            Burny is a burn-up chart generation tool designed to visualize and manage scrum project
                            progress. It makes easy to predict completion dates and coordinate project scopes. Complex
                            project management made simple with an intuitive interface.
                        </p>
                    </div>
                    <div class="chart-container">
                        <Line v-if="chartReady" :data="burnUpChartData" :options="burnUpChartOptions" />
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
                        <h3>Project Tracking</h3>
                        <p>Centrally manage multiple projects and check progress at a glance.</p>
                    </div>
                    <div class="feature-card">
                        <h3>Sprint Management</h3>
                        <p>Effectively manage agile development sprints and improve team productivity.</p>
                    </div>
                </div>
            </div>
        </section>

        <section class="why-section">
            <div class="content-container">
                <h2>Why I Built This</h2>
                <p>
                    I have developed burn-up chart tools in my spare time at two organizations. Because both
                    organizations customized task management systems, their native chart generation capabilities were
                    insufficient. I want to offer an easy-to-use, high-quality solution for burn-up chart management by
                    Burny, enabling development teams to focus on their core business.
                </p>
            </div>
        </section>

        <footer class="site-footer">
            <div class="cta-section">
                <div class="content-container">
                    <h2>Get Started Now</h2>
                    <div class="cta-buttons">
                        <button @click="openSignInModal(false)" class="button-large">Sign In</button>
                        <button @click="openSignInModal(true)" class="button-large">Sign Up</button>
                    </div>
                </div>
            </div>
            <div class="footer-content">
                <div class="content-container">
                    <div class="footer-links">
                        <a href="https://github.com/tomoki108/burny" target="_blank" rel="noopener noreferrer"
                            class="github-link">
                            <span>View Source Code on GitHub</span>
                        </a>
                    </div>
                    <p class="copyright">&copy; {{ new Date().getFullYear() }} Burny. All rights reserved.</p>
                </div>
            </div>
        </footer>

        <SignInModal :isVisible="showSignInModal" :initialSignUp="isSignUp" @close="closeSignInModal"
            @auth-success="handleAuthSuccess" />
    </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';
import { Line } from 'vue-chartjs';
import SignInModal from '../components/SignInModal.vue';

// Register Chart.js components
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const route = useRoute();
const chartReady = ref(false);
const showSignInModal = ref(false);
const isSignUp = ref(false);

onMounted(() => {
    // Delay chart initialization to prevent rendering issues
    setTimeout(() => {
        chartReady.value = true;
    }, 100);

    // Check if redirected from a protected route
    if (route.query.auth === 'required') {
        openSignInModal(false);
    }

    // Check if signup parameter was provided
    if (route.query.signup === 'true') {
        openSignInModal(true);
    }
});

const scrollToAbout = () => {
    document.getElementById('about')?.scrollIntoView({ behavior: 'smooth' });
};

const openSignInModal = (signup = false) => {
    isSignUp.value = signup;
    showSignInModal.value = true;
};

const closeSignInModal = () => {
    showSignInModal.value = false;
};

const handleAuthSuccess = () => {
    showSignInModal.value = false;
};

const burnUpChartData = {
    labels: ['Sprint 1', 'Sprint 2', 'Sprint 3', 'Sprint 4', 'Sprint 5', 'Sprint 6'],
    datasets: [
        {
            label: 'Actual Progress',
            borderColor: '#2196F3',
            backgroundColor: 'rgba(33, 150, 243, 0.1)',
            data: [30, 45, 53, 68, 98, null],
            fill: true,
        },
        {
            label: 'Ideal Progress',
            borderColor: '#FFC107',
            data: [18, 36, 54, 72, 90, 108],
            fill: false,
        },
        {
            label: 'Target Scope',
            borderColor: '#4CAF50',
            data: [110, 110, 110, 110, 110, 110],
            fill: false,
            borderDash: [10, 5],
        },
    ],
};

const burnUpChartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    scales: {
        y: {
            beginAtZero: true,
            title: {
                display: true,
                text: 'Story Points'
            }
        },
        x: {
            title: {
                display: true,
                text: 'Sprint'
            }
        }
    }
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
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    background: linear-gradient(var(--color-tertiary), var(--color-secondary));
    color: var(--color-text-light);
    padding: 2rem 1rem;
    box-sizing: border-box;
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
    padding: 0 2rem;
    width: 100%;
    box-sizing: border-box;
}

.about-section,
.features-section,
.why-section,
.cta-section {
    padding: 3rem 0;
    width: 100%;
    box-sizing: border-box;
    overflow: hidden;
}

.about-section {
    padding-top: 3rem 0;
    padding-bottom: 0;

}

.about-content {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    margin-top: 0.5rem;
}

.about-text {
    width: 100%;
}

.chart-container {
    width: 100%;
    height: 300px;
    position: relative;
    margin-top: 1rem;
}

.features-section {
    background-color: var(--color-background-alt);
}

.features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.5rem;
    margin-top: 1.5rem;
}

.feature-card {
    background-color: var(--color-background);
    padding: 1.5rem;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.why-section {
    background-color: var(--color-background);
}

.cta-section {
    padding: 3rem 0 2rem;
    text-align: center;
}

.footer-content {
    padding: 2rem 0 1.5rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.cta-buttons {
    display: flex;
    justify-content: center;
    gap: 1rem;
    margin-top: 1.5rem;
}

.site-footer {
    background: var(--color-secondary-tertiary);
    color: var(--color-text-light);
    width: 100%;
}

.footer-links {
    display: flex;
    justify-content: center;
    margin-bottom: 0.75rem;
}

.github-link {
    color: var(--color-text-light);
    text-decoration: none;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-weight: bold;
}

.github-link:hover {
    color: var(--color-primary);
    text-decoration: underline;
}

.copyright {
    text-align: center;
    font-size: 0.875rem;
    margin-top: 0.75rem;
    color: rgba(255, 255, 255, 0.7);
}

h2 {
    margin-top: 0;
    margin-bottom: 0.75rem;
}

h3 {
    margin-top: 0;
    margin-bottom: 0.5rem;
}

/* Responsive adjustments */
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

    .hero-actions {
        flex-direction: column;
        width: 100%;
        max-width: 300px;
    }

    .cta-buttons {
        flex-direction: column;
        width: 100%;
        max-width: 300px;
        margin: 1.5rem auto 0;
    }

    .chart-container {
        width: 100%;
        margin-top: 2rem;
    }

    .about-section,
    .features-section,
    .why-section,
    .cta-section {
        padding: 2rem 0;
    }

    .hero-actions,
    .cta-buttons {
        flex-direction: column;
        width: 100%;
        max-width: 200px;
        margin-left: auto;
        margin-right: auto;
    }

    .hero-actions .button,
    .cta-buttons .button {
        width: 100%;
    }
}

/* For very small screens */
@media (max-height: 600px) {
    .hero {
        min-height: 500px;
        padding: 2rem 1rem;
    }
}

@media (max-width: 480px) {
    .hero-title {
        font-size: 2.5rem;
    }

    .feature-card {
        padding: 1.5rem;
    }
}
</style>