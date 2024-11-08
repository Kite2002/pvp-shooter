import { BlendMode, createPicture, Skia } from "@shopify/react-native-skia";
import { SharedValue } from "react-native-reanimated";

type ProjectileType = {
  x: SharedValue<number>;
  y: SharedValue<number>;
  radius: number;
  color: string;
  velocity: {
    x: number;
    y: number;
  };
};

export class Projectile {
  x: SharedValue<number>;
  y: SharedValue<number>;
  radius: number;
  color: string;
  velocity: {
    x: number;
    y: number;
  };

  constructor({ x, y, radius, color, velocity }: ProjectileType) {
    this.x = x;
    this.y = y;
    this.radius = radius;
    this.color = color;
    this.velocity = velocity;
  }
  //   draw() {
  //     return createPicture((canvas) => {
  //       const size = 256;
  //       const r = 0.33 * size;
  //       const paint = Skia.Paint();
  //       paint.setBlendMode(BlendMode.Multiply);

  //       paint.setColor(Skia.Color("cyan"));
  //       canvas.drawCircle(r, r, r, paint);

  //       paint.setColor(Skia.Color("magenta"));
  //       canvas.drawCircle(size - r, r, r, paint);

  //       paint.setColor(Skia.Color("yellow"));
  //       canvas.drawCircle(size / 2, size - r, r, paint);
  //     });
  //   }
}
