import React, { useEffect, useMemo, useRef, useState } from "react";
import { Canvas, Circle, SkiaDomView } from "@shopify/react-native-skia";
import { Dimensions } from "react-native";
import { Player } from "./src/Entities/Player";
import { Projectile } from "./src/Entities/Projectile";
import { makeMutable, withTiming } from "react-native-reanimated";
const App = () => {
  let canvasRef = useRef<SkiaDomView>(null);
  const Window = Dimensions.get("window");
  const cWidth = Window.width;
  const cHeight = Window.height;
  const [projectiles, setProjectiles] = useState<Projectile[]>([]);
  const [player, setPlayer] = useState<Player>(
    new Player({
      x: makeMutable(cWidth / 2),
      y: makeMutable(cHeight / 2),
      radius: 30,
      color: "red",
    })
  );

  function drawPlayer() {
    return (
      <Circle
        cx={player.x}
        cy={player.y}
        r={player.radius}
        color={player.color}
      />
    );
  }

  function drawProjectiles() {
    return (
      <>
        {projectiles.map((projectile, index) => (
          <Circle
            key={index}
            cx={projectile.x}
            cy={projectile.y}
            r={projectile.radius}
            color={projectile.color}
          />
        ))}
      </>
    );
  }

  return (
    <Canvas
      onTouch={(e) => {
        if (e[0][0].type == 2) {
          const x = makeMutable(e[0][0].x);
          const y = makeMutable(e[0][0].y);
          const newProjectile = new Projectile({
            x,
            y,
            radius: 10,
            color: "blue",
            velocity: { x: 5, y: 5 },
          });
          player.x.value = withTiming(e[0][0].x);
          player.y.value = withTiming(e[0][0].y);
          setProjectiles((o) => [...o, newProjectile]);
        }
      }}
      ref={canvasRef}
      style={{ height: cHeight, width: cWidth }}
    >
      {drawPlayer()}
      {drawProjectiles()}
    </Canvas>
  );
};
export default App;
