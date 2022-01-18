import instance from './axios';

// 测试
interface VisualMeasureNum {
  start: number;
  num: number;
}
export const getVisualMeasureNum = (param: VisualMeasureNum) => {
  return instance.post('/visualmeasureNum', param);
};
