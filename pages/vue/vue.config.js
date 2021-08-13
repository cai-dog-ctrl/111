const port = 8589; // dev port
module.exports = {
    devServer: {
        port,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        },
        proxy: {
            "/": {
                target: "http://127.0.0.1:8002",
                changeOrigin: true, // 是否改变域名
                ws: true,
                pathRewrite: {
                    // 路径重写
                    "/": ""
                }
            }
        }

    }
}
