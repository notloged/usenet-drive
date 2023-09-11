import DiskUsageCard, { DiskUsageCardProps } from './DiskUsageCard';

export default {
    component: DiskUsageCard,
    title: 'DiskUsageCard',
    tags: ['autodocs'],
};

export const Default: { args: DiskUsageCardProps } = {
    args: {
        data: {
            folder: '/some/folder',
            total: 1073741824,
            used: 1024,
            free: 1073740800,
        },
    },
};

export const LargeSize = {
    args: {
        data: {
            folder: '/some/folder',
            total: 1099511627776,
            used: 1073741824 * 50,
            free: 1099511627776 - (1073741824 * 50),
        },
    },
};