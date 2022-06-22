import create from 'zustand';

const useStore = create((set) => ({
    dollars: 0,
    broke: false,
    increaseDollars: () => set((state) => ({ dollars: state.dollars + 1 })),
    decreaseDollars: () => set((state) => ({ dollars: state.dollars - 1 })),
    setBroke: (input) => set((state) => ({ broke: input })),
}));