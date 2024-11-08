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
  const [player] = useState<Player>(
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

  let timeout: undefined | NodeJS.Timeout;
  useEffect(() => {
    projectiles.length > 0 && console.log(projectiles);
    timeout && clearTimeout(timeout);
    timeout = setInterval(() => {
      projectiles.forEach((p) => {
        p.x.value += p.velocity.x;
        p.y.value += p.velocity.y;
      });
      setProjectiles((oProjectiles) => {
        const updatedProjectiles = oProjectiles.filter(
          (p) =>
            p.x.value >= 0 &&
            p.x.value < cWidth &&
            p.y.value >= 0 &&
            p.y.value < cHeight
        );
        // console.log(updatedProjectiles);
        updatedProjectiles.forEach((p) => {
          p.x.value += p.velocity.x;
          p.y.value += p.velocity.y;
        });
        return updatedProjectiles;
      });
    }, 16);
    return () => timeout && clearTimeout(timeout);
  }, [projectiles]);

  return (
    <Canvas
      onTouch={(e) => {
        if (e[0][0].type == 2) {
          // const x = makeMutable(e[0][0].x);
          // const y = makeMutable(e[0][0].y);
          const angle = Math.atan2(
            e[0][0].y - cHeight / 2,
            e[0][0].x - cWidth / 2
          );
          const x = makeMutable(cWidth / 2);
          const y = makeMutable(cHeight / 2);
          const newProjectile = new Projectile({
            x,
            y,
            radius: 10,
            color: "blue",
            velocity: { x: Math.cos(angle) * 10, y: Math.sin(angle) * 10 },
          });
          // player.x.value = withTiming(e[0][0].x);
          // player.y.value = withTiming(e[0][0].y);
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
