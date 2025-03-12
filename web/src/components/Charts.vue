<template>
    <v-tabs v-model="activeTab">
        <v-tab value="burn">Burn Up Chart</v-tab>
        <v-tab value="velocity">Velocity Chart</v-tab>
    </v-tabs>
    <v-tabs-window v-model="activeTab">
        <v-tabs-window-item value="burn">
            <Line class="pr-16" :data="burnUpChartData" :options="burnUpchartOptions" />
        </v-tabs-window-item>
        <v-tabs-window-item value="velocity">
            <Line class="pr-16" :data="velocityChartData" :options="velocityChartOptions" />
        </v-tabs-window-item>
    </v-tabs-window>
</template>

<script setup lang="ts">
import { type Sprint } from '../api/sprint_api';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale } from 'chart.js';
import { Line } from 'vue-chartjs';
import { computed, ref } from 'vue';
import { isSprintStarted } from '../utils/sprint_helper';

const props = defineProps<{
    sprints: Sprint[]
    total_sp: number
}>();

const activeTab = ref(0);

ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale);

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
        scales: {
            y: {
                max: max,
                ticks: {
                    stepSize: 10,
                },
            },
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
            label: 'velocity',
            borderColor: '#ff5252', // var(--color-danger)
            data: velocity.value,
            fill: true,
        },
    ],
}));

const velocityChartOptions = computed(() => {
    return {
        responsive: true,
        scales: {
            y: {
                min: 0,
            },
        },
    };
});



</script>