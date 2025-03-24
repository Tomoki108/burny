<template>
    <v-tabs v-model="activeTab">
        <v-tab value="burn">Burn-Up Chart</v-tab>
        <v-tab value="velocity">Velocity Chart</v-tab>
    </v-tabs>
    <v-tabs-window v-model="activeTab">
        <v-tabs-window-item value="burn">
            <div class="chart-container">
                <Line :data="burnUpChartData" :options="burnUpchartOptions" />
            </div>
        </v-tabs-window-item>
        <v-tabs-window-item value="velocity">
            <div class="chart-container">
                <Line :data="velocityChartData" :options="velocityChartOptions" />
            </div>
        </v-tabs-window-item>
    </v-tabs-window>
</template>

<script setup lang="ts">
import { type Sprint } from '../api/sprint_api';
import { Line } from 'vue-chartjs';
import { computed, ref } from 'vue';
import { isSprintStarted } from '../utils/sprint_helper';

const props = defineProps<{
    sprints: Sprint[]
    total_sp: number
}>();

const activeTab = ref(0);
const windowWidth = ref(window.innerWidth);

// Burn Up Chart
const cumulativeActualSp = computed(() => {
    const sprints = props.sprints.filter(sprint => isSprintStarted(sprint));
    let cumulative = 0;
    return sprints.map(sprint => {
        cumulative += sprint.actual_sp;
        return cumulative;
    });
});

const cumulativeIdealSp = computed(() => {
    const sprints = props.sprints;
    let cumulative = 0;
    return sprints.map(sprint => {
        cumulative += sprint.ideal_sp;
        return cumulative;
    });
});

const burnUpChartData = computed(() => ({
    labels: props.sprints.map((_, index) => `Sprint ${index + 1}`),
    datasets: [
        {
            label: 'Actual Progeress',
            borderColor: '#2196f3', // var(--color-info)
            data: cumulativeActualSp.value,
            fill: true,
        },
        {
            label: 'Ideal Progress',
            borderColor: '#ffc107', // var(--color-warning)
            data: cumulativeIdealSp.value,
            fill: false,
        },
        {
            label: 'Target Scope',
            borderColor: '#4caf50', // var(--color-success)
            data: Array(props.sprints.length).fill(props.total_sp),
            fill: false,
            borderDash: [10, 5],
        },
    ],
}));

const burnUpchartOptions = computed(() => {
    let max = props.total_sp + 10
    if (props.total_sp % 10 !== 0) {
        max -= props.total_sp % 10
    }

    return {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
            y: {
                max: max,
                ticks: {
                    stepSize: 10,
                },
                title: {
                    display: windowWidth.value > 390,
                    text: 'Story Points'
                }
            },
            x: {
                title: {
                    display: windowWidth.value > 390,
                    text: 'Sprints'
                }
            }
        },
    };
});

// Velocity Chart
const velocity = computed(() => {
    // 既に開始日を過ぎているスプリントのactyual_spを取得
    return props.sprints.filter(sprint => isSprintStarted(sprint)).map(sprint => sprint.actual_sp);
});

const velocityChartData = computed(() => ({
    labels: props.sprints.map((_, index) => `Sprint ${index + 1}`),
    datasets: [
        {
            label: 'Velocity',
            borderColor: '#ff5252', // var(--color-danger)
            data: velocity.value,
            fill: true,
        },
    ],
}));

const velocityChartOptions = computed(() => {
    return {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
            y: {
                min: 0,
                title: {
                    display: windowWidth.value > 390,
                    text: 'Story Points'
                }
            },
            x: {
                title: {
                    display: windowWidth.value > 390,
                    text: 'Sprints'
                }
            }
        },
    };
});
</script>

<style scoped>
.chart-container {
    position: relative;
    height: 650px;
    width: 100%;
}
</style>