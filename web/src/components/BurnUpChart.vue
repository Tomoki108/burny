<template>
    <Line class="pr-16" :data="chartData" :options="chartOptions" />
</template>

<script setup lang="ts">
import { type Sprint } from '../api/sprint_api';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale } from 'chart.js';
import { Line } from 'vue-chartjs';
import { computed } from 'vue';

ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale);

const props = defineProps<{
    sprints: Sprint[]
    total_sp: number
}>()

// Calculate cumulative actual_sp
const cumulativeActualSp = computed(() => {
    const sprints = props.sprints;
    let cumulative = 0;
    return sprints.map(sprint => {
        cumulative += sprint.actual_sp;
        return cumulative;
    });
});

// Calculate cumulative ideal_sp
const cumulativeIdealSp = computed(() => {
    const sprints = props.sprints;
    let cumulative = 0;
    return sprints.map(sprint => {
        cumulative += sprint.ideal_sp;
        return cumulative;
    });
});

const chartData = computed(() => ({
    labels: props.sprints.map((_, index) => `Sprint ${index + 1}`),
    datasets: [
        {
            label: 'accumulated actual_sp',
            borderColor: '#2196f3',// var(--color-info)
            data: cumulativeActualSp.value,
            fill: true,
        },
        {
            label: 'accumulated ideal_sp',
            borderColor: '#ffc107', // var(--color-warning)
            data: cumulativeIdealSp.value,
            fill: false,
        },
        {
            label: 'total_sp',
            borderColor: '#4caf50', // var(--color-success)
            data: Array(props.sprints.length).fill(props.total_sp),
            fill: false,
            borderDash: [10, 5],
        },
    ],
}));

const chartOptions = computed(() => {
    return {
        responsive: true,
        scales: {
            y: {
                max: props.total_sp + 10,
                ticks: {
                    stepSize: 10,
                },
            }
        }
    }
});
</script>