import service from '@/utils/request'

export const getRecolistList = (params) => {
    return service({
        url: "/recolist/getRecolistList",
        method: 'get',
        params
    })
}