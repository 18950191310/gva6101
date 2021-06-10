import { login } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            userName: '',
            nickName: "",
            headerImg: "",
            authority: "",
        },
        oderDetail:{},
        total: "",
        token: "",
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo

        },
        setOrderDetail(state, orderDetail) {
            state.orderDetail = orderDetail
        },
        setTotal(state, Total) {
            state.total = Total
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        },
        LoginOut(state) {
            state.userInfo = {}
            state.token = ""
            sessionStorage.clear()
            router.push({ name: 'login', replace: true })
            window.location.reload()
        },
        ResetUserInfo(state, userInfo = {}) {
            state.userInfo = {...state.userInfo,
                ...userInfo
            }
        }
    },

    actions: {
        async LoginIn({ commit, dispatch, rootGetters, getters }, loginInfo) {
            const res = await login(loginInfo)
            if (res.code == 0) {
                commit('setUserInfo', res.data.user)
                commit('setToken', res.data.token)
                await dispatch('router/SetAsyncRouter', {}, { root: true })
                const asyncRouters = rootGetters['router/asyncRouters']
                asyncRouters.map(asyncRouter => {
                    router.addRoute(asyncRouter)
                })
                // const redirect = router.currentRoute._value.query.redirect
                // if (redirect) {
                //     router.push({ path: redirect })
                // } else {
                router.push({ name: getters["userInfo"].authority.defaultRouter })
                // }
                return true
            }
        },
        async LoginOut({ commit }) {
            const res = await jsonInBlacklist()
            if (res.code == 0) {
                commit("LoginOut")
            }
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },
        orderDetail(state) {
            return state.orderDetail
        },
        total(state) {
           return state.total
        }

    }
}