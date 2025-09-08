export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig();
    const method = event.node.req.method?.toLocaleUpperCase()
})