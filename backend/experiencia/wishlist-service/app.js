import express from "express";
import dotenv from "dotenv";
import wishlistRoutes from "./routes/wishlistRoutes.js";

dotenv.config();
const app = express();
app.use(express.json());
app.use("/", wishlistRoutes);

const PORT = process.env.PORT || 3026;
app.listen(PORT, () => {
  console.log(`wishlist-service corriendo en el puerto ${PORT}`);
});
