import express from "express";
import { obtenerWishlist } from "../controllers/getWishlist.js";
import { agregarAWishlist } from "../controllers/addToWishlist.js";
import { eliminarDeWishlist } from "../controllers/removeFromWishlist.js";

const router = express.Router();

router.get("/wishlist", obtenerWishlist);
router.post("/wishlist", agregarAWishlist);
router.delete("/wishlist/:productId", eliminarDeWishlist);

export default router;
