import LOGIN from '../view/LoginModule/LoginPage'
import loginView from '../view/LoginModule/LoginView'
import registerView from '../view/LoginModule/RegisterView'
import VOTE from '../view/VoteModule/VotePage'
import CREATE from '../view/VoteModule/CreatePage'
import VOTING from '../view/VoteModule/VotingPage'
import {createRouter,createWebHistory} from 'vue-router'

const routes=[
    {
        path: "/",
        redirect:"/login"
    },
    {
        path:"/",
        component:LOGIN,
        children:[
            {path:"login",component:loginView},
            {path:"register",component:registerView}
        ]
    },
    {
        path:"/vote",
        component: VOTE
    },
    {
        path:"/vote/create",
        component: CREATE,
    },
    {
        path: "/vote/voting/:addr",
        component: VOTING
    }
]

const router=createRouter({
    history:createWebHistory(),
    routes
})

export default router