import create from 'zustand';

const useStore = create((set) => ({
    token: '',
    id: '',
    name: '',
    role: '',
    setToken: (token) => set(() => ({ token })),
    setId: (id) => set(() => ({ id })),
    setName: (name) => set(() => ({ name })),
    setRole: (role) => set(() => ({ role })),
}));

export default useStore;
export { useStore };