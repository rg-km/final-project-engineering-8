import create from 'zustand';

const useStore = create((set) => ({
    token: '',
    id: '',
    name: '',
    setToken: (token) => set(() => ({ token })),
    setId: (id) => set(() => ({ id })),
    setName: (name) => set(() => ({ name })),
}));

export default useStore;
export { useStore };