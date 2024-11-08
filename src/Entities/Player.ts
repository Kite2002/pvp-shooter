import { SharedValue } from "react-native-reanimated";

type PlayerType = {
  x: SharedValue<number>;
  y: SharedValue<number>;
  radius: number;
  color: string;
};

export class Player {
  x: SharedValue<number>;
  y: SharedValue<number>;
  radius: number;
  color: string;

  constructor({ x, y, radius, color }: PlayerType) {
    this.x = x;
    this.y = y;
    this.radius = radius;
    this.color = color;
  }
}
