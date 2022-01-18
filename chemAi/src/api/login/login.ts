import instance from '../axios';

// 登录
interface LoginParam {
  password: string;
  username: string;
}
export const loginApi = (param: LoginParam) => {
  return instance.post('/user/login', param);
};

// 修改密码
interface ChangePasswordParam {
  newPassword: string;
  oldPassword: string;
  repeatPassword: string;
  username: string;
}
export const changePasswordApi = (param: ChangePasswordParam) => {
  return instance.post('/user/changePassword', param);
};
