<template>
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
</template>

<script lang="ts" setup>
const emit = defineEmits(['open-signin-modal']);

const openSignInModal = (isSignUp: boolean) => {
    emit('open-signin-modal', isSignUp);
};
</script>

<style scoped>
.site-footer {
    background: var(--color-secondary-tertiary);
    color: var(--color-text-light);
    width: 100%;
}

.cta-section {
    padding: 3rem 0 2rem;
    text-align: center;
}

.footer-content {
    padding: 2rem 0 1.5rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.content-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
    width: 100%;
    box-sizing: border-box;
}

.cta-buttons {
    display: flex;
    justify-content: center;
    gap: 1rem;
    margin-top: 1.5rem;
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

/* Responsive adjustments */
@media (max-width: 768px) {
    .cta-buttons {
        flex-direction: column;
        width: 100%;
        max-width: 300px;
        margin: 1.5rem auto 0;
    }

    .cta-section {
        padding: 2rem 0;
    }

    .cta-buttons {
        flex-direction: column;
        width: 100%;
        max-width: 200px;
        margin-left: auto;
        margin-right: auto;
    }

    .cta-buttons .button {
        width: 100%;
    }
}
</style>