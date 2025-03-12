<template>
    <div class="home-container">
        <header class="hero">
            <h1 class="hero-title">Burny 🐶</h1>
            <p class="hero-subtitle">Simplify Project Management with Burn-up Charts</p>
            <div class="hero-actions">
                <button @click="openSignInModal(false)" class="button primary">Sign In</button>
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
                <h2>Why We Built This</h2>
                <p>
                    This project began with the desire to manage team progress more efficiently.
                    By providing features missing in existing tools and a more intuitive interface,
                    we aim to reduce the burden of project management and allow teams to focus on higher-value work.
                </p>
            </div>
        </section>

        <footer class="site-footer">
            <div class="cta-section">
                <div class="content-container">
                    <h2>Get Started Now</h2>
                    <div class="cta-buttons">
                        <button @click="openSignInModal(false)" class="button primary">Sign In</button>
                        <button @click="openSignInModal(true)" class="button primary">Sign Up</button>
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

        <!-- Sign In Modal -->
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
    /* Reduced padding */
    box-sizing: border-box;
    /* Include padding in height calculation */
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
    /* Reduced padding from 5rem to 3rem */
    width: 100%;
    box-sizing: border-box;
    overflow: hidden;
    /* Prevent content overflow */
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
    /* Reduced gap */
    margin-top: 1.5rem;
    /* Reduced margin */
}

.feature-card {
    background-color: var(--color-background);
    padding: 1.5rem;
    /* Reduced padding */
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
    /* Reduced margin */
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

.button {
    font-family: var(--font-family-base);
    background-color: var(--color-primary);
    color: var(--color-text-light) !important;
    border: none;
    padding: 12px;
    font-size: 16px;
    border-radius: 4px;
    cursor: pointer;
    margin: 5px 5px;
    box-shadow: 0.1rem 0.1rem var(--color-shadow);
    font-weight: 500;
    width: auto;
    /* 自動幅に変更 */
    min-width: 120px;
    /* 最小幅を設定 */
    display: inline-block;
}

.button.primary {
    background-color: var(--color-primary);
    color: var(--color-text-light) !important;
}

.button.secondary {
    background-color: transparent;
    color: var(--color-text-light) !important;
    border: 2px solid var(--color-text-light);
    box-shadow: none;
}

.button.primary:hover {
    background-color: var(--color-primary-dark);
}

.button.secondary:hover {
    background-color: rgba(255, 255, 255, 0.1);
}

/* 見出しのマージンを調整 */
h2 {
    margin-top: 0;
    margin-bottom: 0.75rem;
    /* 1.5remから0.75remに縮小 */
}

h3 {
    margin-top: 0;
    margin-bottom: 0.5rem;
    /* 0.75remから0.5remに縮小 */
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
        /* Reduced margin */
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
        /* Further reduce padding on mobile */
    }

    .hero-actions,
    .cta-buttons {
        flex-direction: column;
        width: 100%;
        max-width: 200px;
        /* ボタン幅を制限 */
        margin-left: auto;
        margin-right: auto;
    }

    .hero-actions .button,
    .cta-buttons .button {
        width: 100%;
        /* モバイルでは幅いっぱいに */
    }
}

/* For very small screens */
@media (max-height: 600px) {
    .hero {
        min-height: 500px;
        padding: 2rem 1rem;
        /* Reduced padding */
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